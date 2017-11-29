<template>
  <div id="app">
    <img src="/assets/logo.png">
    <h1>{{ msg }}</h1>
    <form method="post">
      <label for="login">Login</label>
      <input type="text" name="login" id="login" v-model="user.login" />
      <label for="password">Password</label>
      <input type="password" name="password" id="password" v-model="user.password" />
      <input type="button" value="login" @click="login" />
      <input type="button" value="register" @click="register" />
    </form>
    <br />
    <app-user-page v-if="authenticated" v-bind:token="token" v-bind:userdata="userdata"></app-user-page>
    <div v-if="message !== null">
      {{ message }} : {{ statusCode }}
    </div>
  </div>
</template>

<script>
import UserPage from "./UserPage.vue";
export default {
  name: 'app',
  components: {
    "app-user-page": UserPage
  },
  data () {
    return {
      msg: 'Welcome to Your Vue.js App',
      user:{
        login: "",
        password: ""
      },
      authenticated : false,
      token: null,
      message: null,
      statusCode: null,
      userdata: null
    }
  },
  methods:{
    login: function(){
      this.$http.post("api/login", this.user).then(response => {
        this.token = response.body;
        this.$http.get("api/userdata?login=" + this.user.login, { headers: { Authorization: ("Bearer " + this.token) }}).then(response =>{
          this.authenticated = true;
          this.userdata = response.body;
          this.message = null;
          this.statusCode = null;
        }, error => {
          this.authenticated = false;
          this.token = null;
          this.message = error.body;
          this.statusCode = error.status;
        });
      }, error => {
        this.authenticated = false;
        this.token = null;
        this.message = error.body;
        this.statusCode = error.status;
      });
    },
    register: function(){
      this.$http.post("api/users/create", this.user).then(response => function(){
        this.message = "Registration Successful";
        this.statusCode = response.statusCode;
      }, error => {
        this.message = error.body;
        this.statusCode = error.status;
      });
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}

h1, h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}
</style>
