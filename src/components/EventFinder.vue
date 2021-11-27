<template>
  <div class="search">
    <h1>{{ msg }}</h1>
    <form v-on:submit.prevent="findEvents">
      <div class="form-group col-sm-2 offset-5 center-block">
        <select class="form-control" v-model="genreInput" id="genre-input">
          <option selected disabled>Genre</option>
          <option>Trance</option>
          <option>House</option>
        </select>
      </div>
      <div class="form-group col-sm-2 offset-5 center-block">
        <select class="form-control" v-model="locationInput" id="location-input">
          <option selected disabled>Location</option>
          <option>San Diego</option>
          <option>Los Angeles</option>
        </select>
      </div>
      <div class="form-group">
        <button class="btn btn-light"><img alt="MusicSnob" src="../assets/musicnote.png" width="20" height="20"></button>
      </div>
    </form>
    <ul class="list-group">
      <li v-for="item in items" class="list-element" v-bind:key="item.id">
        <div class="row align-items-start">
          <div class="col-sm-4 offset-4">
            {{ item.date }}
            <br>
            {{ item.artist }} 
            <br>
            {{ item.venue }}
            <br><br>
          </div>
        </div>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'EventFinder',
  props: {
    msg: String
  },

  data() { return {
    genreInput: '',
    locationInput: '',
    items: [],
  } },

  methods: {
    findEvents() {
      console.log(`Search for ${this.genreInput} events in ${this.locationInput}`);
      var data = {"genre": this.genreInput, "location": this.locationInput}
      axios({ method: "POST", url: "http://localhost:3000/api", data: data, headers: {"content-type": "text/plain" } }).then(result => { 
          console.log(result.data) 
          this.items = result.data
        }).catch( error => {
            console.error(error);
      });
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
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
