<template>
    <div id="app">

        <TweetArea/>
    </div>
</template>

<script>

    import TweetArea from "./components/TweetArea"


    export default {
        name: 'App',
        components: {
            TweetArea
        }
    }
</script>

<style>
    #app {
        font-family: Avenir, Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
        text-align: center;
        color: #2c3e50;
        margin-top: 60px;
    }
</style>
