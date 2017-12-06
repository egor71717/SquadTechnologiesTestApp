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
                    console.log("err");
                    this.message = error.body;
                    this.statusCode = error.status;
                });
            }, error => {
                this.message = error.body;
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
            location.href="auth/getcode/vk";
            // this.$http.get("auth/getcode/vk").then(response => {
            //     let authData = response.body
            //     this.message = "your access token: " + authData.access_token + "\nuser ID: " + authData.user_id;
            //     this.statusCode = null;
            // }, error => {
            //     this.message = error.body;
            //     this.statusCode = error.status;
            // });
        },
        loginfacebook: function(){
            //https://www.facebook.com/v2.11/dialog/oauth?client_id=134698900578486&redirect_uri=http://localhost:8800

            const FBAppID  = "134698900578486"

            this.$http.get("https://www.facebook.com/dialog/oauth", {
                client_id: FBAppID,
                redirect_uri: "http://localhost:8800/auth/facebook.html"
            }).then(response => {
                console.log(response.body);
                console.log(response.statusCode);
            }, error => {
                console.log(response.body);
                console.log(response.statusCode);
            });
        }
  }
}
</script>
