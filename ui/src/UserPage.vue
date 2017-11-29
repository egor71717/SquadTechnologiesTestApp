<template>
    <div id="app-user-page">
        Authenticated successfully, your token:{{ token }}
        <p>login: {{ userdata.Login }}</p>
        <p>password hash: {{ userdata.PasswordHash }}</p>
    </div>
</template>

<script>
export default {
    name: "app-user-page",
    props:{
        token:{
            type: String,
            required: true
        },
        userdata:{
            type: Object
        }
    }
}
</script>
