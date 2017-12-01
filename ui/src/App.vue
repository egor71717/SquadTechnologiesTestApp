<template>
  <div id="app">
    <img src="/assets/logo.png">
    <h1>
      {{ message }} 
      <span v-if="statusCode">:{{ statusCode }}</span> 
    </h1>
    <br />
    <div v-if="!authenticated">
      <app-authentication @authenticated="onAuthenticate($event)"></app-authentication>
    </div>
    <div v-if="authenticated">
      <app-user-page  v-bind:token="token" v-bind:userdata="userdata"></app-user-page>
      <input type="button" value="logout" @click="logout"/>
    </div>
  </div>
</template>
<script>
import UserPage from "./UserPage.vue";
import Authentication from "./Authentication.vue"
export default {
  name: 'app',
  components: {
    "app-user-page": UserPage,
    "app-authentication": Authentication
  },
  data () {
    return {
      message: 'Welcome to Your Vue.js App',
      authenticated : false,
      token: null,
      statusCode: null,
    }
  },
  methods:{
    logout: function(){
      this.authenticated = false;
      this.token = false;
    },
    onAuthenticate: function(event){
      this.authenticated = true;
      this.token = event.token;
      this.userdata = event.userdata;
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
