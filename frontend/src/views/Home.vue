<template>
  <v-app>
    <v-app-bar app color="primary" dark> </v-app-bar>
    <v-container>
      <v-main>
        <v-row>
          <v-text-field v-model="name" label="user" required></v-text-field>
        </v-row>
        <v-row>
          <v-text-field
            v-model="server"
            label="Chat server"
            required
          ></v-text-field>
        </v-row>
        <v-btn 
        v-on:click="this.connectServer"
        depressed color="primary">
          Connect
        </v-btn>
      </v-main>

      <v-card
        height="100%"
        class="mx-auto mt-8"
        color="#F9EDEB"
        dark
        min-width="1000"
      >
       <li v-for="item in get" :key="item.user">
    {{ item.user }} -- {{item.message}}
  </li>
      </v-card>
      <v-row>
        <v-text-field
          class="mt-10"
          v-model="message"
          label="Message"
        ></v-text-field>
      </v-row>
      <v-row>
        <v-btn 
        v-on:click="this.sendMessage"
        depressed 
        color="info">
          Send
        </v-btn>
      </v-row>
    </v-container>
    <v-container>
      <!-- <v-row>
          <v-text-field v-model="message" label="Message"></v-text-field>
        </v-row> -->
    </v-container>
  </v-app>
</template>
<script>
import axios from 'axios';
const evtSource = new EventSource("https://chat-room-be.herokuapp.com", { withCredentials: true } );
export default {
  name: "Home",
  components: {
    // HelloWorld,
  },
  data() {
    return {
      name: "brown",
      server: "https://chat-room-be.herokuapp.com",
      message: "",
      get: [],
    };
  },
  created() {
    const hash = Math.floor(Math.random() * 1000000) + 1000;
    console.log(hash);
    this.name = this.name + hash;
  },
  mounted() {
    axios.get(`http://localhost:9999/chat/message`)
    .then(response => {
      console.log("res: ")
      console.log(response.data.data)
      this.get =response.data.data
      // console.log(this.posts)
    })
    .catch(e => {
      this.errors.push(e)
    })
  },
  methods: {
    sendMessage() {
      console.log(this.message)
      axios.post(`http://localhost:9999/chat/message`, {
    user: this.name,
    message: this.message
  })
  .then(function (response) {
    console.log(response);
  })
  .catch(function (error) {
    console.log(error);
  });
    },
    connectServer() {
      console.log("connect to server")
      console.log(evtSource)
    }
  },
  watch: {
    posts() {
      console.log("--");
      console.log(this.posts.length);
      let i = 0;
      for (i = 0; i < this.posts.length; i++) {
        console.log(this.posts[i].Age);
      }
    },
  },
};
</script>
