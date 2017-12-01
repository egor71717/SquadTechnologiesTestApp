<template>
    <div>
        <form v-if="!authenticated" >
            <label for="login">Login</label>
            <input type="text" name="login" id="login" v-model="user.login" />
            <label for="password">Password</label>
            <input type="password" name="password" id="password" v-model="user.password" />
            <input type="button" value="login" @click="login" />
            <input type="button" value="register" @click="register" />
            <input type="button" value="login vk" @click="loginvk"/>
            <input type="button" value="login facebook" @click="loginfacebook"/>
        </form>
        <p v-if="message">{{ message }}</p>
        <p v-if="statusCode">{{ statusCode }}</p> 
    </div>
</template>

<script>
export default {
    data(){
        return {
            message: null,
            statusCode: null,
            user:{
                login: "",
                password: ""
            },
            userdata: null,
            token: null
        }
    },
    methods:{
        login: function(){
            this.$http.post("api/login", this.user).then(response => {
                this.token = response.body;
                this.$http.get("api/userdata?login=" + this.user.login, { headers: { Authorization: ("Bearer " + this.token) }}).then(response => {
                    this.message = null;
                    this.statusCode = null;
                    this.$emit("authenticated",{ token: this.token, userdata: response.body});
                }, error => {
                    this.message= error.body;
                    this.statusCode = error.status;
                });
            }, error => {
                this.message= error.body;
                this.statusCode = error.status;
            });
        },
        register: function(){
            if(this.user.password === "" || this.user.login === ""){
                this.message = "invalid user input";
                this.statusCode = null;
                return;
            }
            this.$http.post("api/users/create", this.user).then(response => {
                this.message = "Registration Successful";
                this.statusCode = response.statusCode;
            }, error => {
                this.message = error.body;
                this.statusCode = error.status;
            });
        },
        loginvk: function(){
            console.error("Vk auth is not implemented");
            return;

            let scopeBits={
                email: 4194304
            }
            //https://oauth.vk.com/authorize?client_id=6280065?redirect_uri=http://localhost:8800?display=page?response_type=token?v=5.69?scope=4194304?state=email
            //https://oauth.vk.com/authorize?client_id=6280065?redirect_uri=http://localhost:8800
            //https://oauth.vk.com/authorize?client_id=6280065?redirect_uri=https://oauth.vk.com/blank.html 
            
            this.$http.get("https://oauth.vk.com/authorize", {
                client_id: "6280065",
                redirect_uri: "http://localhost:8800",
                display: "page",
                response_type: "token",
                v: "5.69",
                scope: "4194304",
                state: "authentication_state"
            }).then(response => {
                console.log(response.body);
                console.log(response.statusCode);
            }, error => {
                console.log(response.body);
                console.log(response.statusCode);
            });
        },
        loginfacebook: function(){
            console.error("Facebook auth is not implemented")
            return;
        }
  }
}
</script>
