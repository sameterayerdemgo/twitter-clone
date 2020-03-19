<template>

    <a href="https://www.google.com">
        <div>
            <link href="https://fonts.googleapis.com/css?family=Asap" rel="stylesheet">
            <link href="https://fonts.googleapis.com/css?family=Roboto" rel="stylesheet">

            <div class="tweet-wrap">
                <div class="tweet-header">
                    <div class="tweet-header-info">
                        Steve Schoger <span>@Steve Schoger</span><span>. Jun 27
</span>
                        <p>🔥 If you're tired of using outline styles for secondary buttons, a soft solid background
                            based on the text color can be a great alternative.</p>

                    </div>

                </div>


            </div>
        </div>
    </a>
</template>

<script>

    export default {
        components: {},
        name: "Post"
    }
</script>

<style scoped>

    template {

        font-family: 'Asap', sans-serif;
        font-family: 'Roboto', sans-serif;

    }

    a {
        
    }

    .tweet-wrap {
        width: 520px;
        height: 130px;
        background: #fff;
        margin: 0 auto;

        border-radius: 3px;
        padding: 30px;
        border: 1px solid #DDDDDD;

    }

    .tweet-header {
        display: flex;
        align-items: flex-start;
        font-size: 14px;
    }

    .tweet-header-info {
        font-weight: bold;
    }

    .tweet-header-info span {
        color: #657786;
        font-weight: normal;
        margin-left: 5px;
    }

    .tweet-header-info p {
        font-weight: normal;
        margin-top: 5px;

    }

    .tweet-img-wrap {
        padding-left: 60px;
    }

    .tweet-info-counts {
        display: flex;
        margin-left: 60px;
        margin-top: 10px;
    }

    .tweet-info-counts div {
        display: flex;
        margin-right: 20px;
    }

    .tweet-info-counts div svg {
        color: #657786;
        margin-right: 10px;
    }

    @media screen and (max-width: 430px) {
        body {
            padding-left: 20px;
            padding-right: 20px;
        }

        .tweet-header {
            flex-direction: column;
        }

        .tweet-header img {
            margin-bottom: 20px;
        }

        .tweet-header-info p {
            margin-bottom: 30px;
        }

        .tweet-img-wrap {
            padding-left: 0;
        }

        .tweet-info-counts {
            display: flex;
            margin-left: 0;
        }

        .tweet-info-counts div {
            margin-right: 10px;
        }
    }
</style>
