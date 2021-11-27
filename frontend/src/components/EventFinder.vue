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
        <button class="btn btn-dark">Go!</button>
      </div>
    </form>
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
  } },

  methods: {
    findEvents() {
      console.log(`Search for ${this.genreInput} events in ${this.locationInput}`);
      axios.get('http://localhost:3000/').then(resp => {console.log(resp.data)})
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
