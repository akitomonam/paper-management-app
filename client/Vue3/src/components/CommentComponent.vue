<template>
    <div>
        <form>
            <textarea v-model="newComment"></textarea>
            <button @click="submitComment">Submit</button>
        </form>
        <div class="comments">
            <div class="comment" v-for="(comment, index) in comments" :key="index">
                {{ comment }}
            </div>
        </div>
    </div>
</template>
<script>
export default {
    data() {
        return {
            newComment: '',
            comments: []
        }
    },
    methods: {
        submitComment() {
            // add newComment to the list of comments
            this.comments.push(this.newComment);
            this.newComment = '';
        }
    }
}
</script>
<style>
form {
    display: flex;
    flex-direction: column;
}

textarea {
    width: 100%;
    height: 100px;
    margin-bottom: 10px;
    padding: 5px;
}

.comments {
    margin-top: 20px;
}

.comment {
    border: 1px solid #ccc;
    padding: 10px;
    margin-bottom: 10px;
}
</style>
