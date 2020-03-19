<template lang="html">
    <div class="container-fluid">
        <div class="row justify-content-lg-center">
            <ProfilePicture></ProfilePicture>
            <div class="tweet-box">
                <textarea placeholder="'placeholder'"  v-model="message" class="tweet-input"></textarea>
                <br>
                <button @click="sharePost" class="active">Tweet</button>
                <span class="danger"> {{calcuateCharacter}}</span>
            </div>
        </div>
        <div class="row justify-content-lg-center">
            <ProfilePicture/>
            <Post/>
        </div>
    </div>
</template>

<script>
    import ProfilePicture from "./ProfilePicture";
    import Post from "./Post";

    export default {
        placeholder : " " ,
        components: {
            Post,
            ProfilePicture
        },
        data() {
            return {
                message: " "
            }
        },
        computed: {
            calcuateCharacter() {

                return 200 - this.message.length;
            }
        },
        methods: {
            sharePost() {

            }
        }

    };
</script>

<style lang="css" scoped>
    .tweet-box {
        font-size: 18px;
        padding: 10px;
        position: relative;
    }

    .tweet-input {
        border: 1px solid #ccc;
        border-radius: 10px;
        box-sizing: border-box;
        font-family: "Times New Roman", Times, serif;
        font-size: 18px;
        padding: 15px 12px;
        outline: none;
        width: 520px;
        height: 130px;

    }

    span {
        float: right;
        padding: 18px 14px;
    }

    .warning {
        color: #a76c00;
    }

    .danger {
        color: #800000;
    }

    button {
        background-color: #ccc;
        border: 1px solid #eee;
        border-radius: 10px;
        color: #fefefe;
        font-size: 13px;
        float: right;
        margin-top: 10px;
        padding: 10px;
    }

    button:hover {
        background-color: #333;
    }

    button.active {
        background-color: #008cff;
        cursor: pointer;
    }
</style>
