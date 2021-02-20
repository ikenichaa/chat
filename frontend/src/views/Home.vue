<template>
  <v-app>
    <v-app-bar app color="primary" dark> </v-app-bar>
    <v-main>
      <h1>Nicha</h1>
      <h1>Jaa runn na</h1>
      <p>{{this.posts.id}}</p>
      <!-- <img alt="Vue logo" src="../assets/logo.png"> -->
      <!-- <HelloWorld msg="Welcome to Your Vue.js App" /> -->
      <v-btn>
      Normal
    </v-btn>
    <ul id="example-1">
  <li v-for="item in posts" :key="item.ID">
    {{ item.Name }}
  </li>
  <li>
    helloooo
  </li>
</ul>
    </v-main>
    
  </v-app>
</template>
<script>
import axios from 'axios';

export default {
  name: "Home",
  components: {
    // HelloWorld,
  },
  data() {
    return {
      post: "",
      posts: [],
    }
  },
  created() {
    axios.get(`http://localhost:9999/plate/worker/get-all`)
    .then(response => {
      console.log("res: ")
      console.log(response.data.data)
      this.posts =response.data.data
      // console.log(this.posts)
    })
    .catch(e => {
      this.errors.push(e)
    })
  },
  mounted() {
    // this.post = "mount"
    // console.log("I am post: "+this.post)
    // console.log(this.posts)
  },
  watch: {
    posts() {
      console.log("--")
      console.log(this.posts.length)
      let i =0;
      for (i=0;i<this.posts.length;i++){
        console.log(this.posts[i].Age)
      }
    }
  }
  
}
</script>
