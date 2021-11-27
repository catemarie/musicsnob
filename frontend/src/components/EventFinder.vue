<template>
  <div class="search">
    <h1>{{ msg }}</h1>
    <form v-on:submit.prevent="findEvents">
      <h3>Genre</h3>
      <div class="form-group col-sm-4 offset-4 center-block">
        <input v-model="genreInput" type="text" id="genre-input" placeholder="genre" class="form-control text-center">
      </div>
      <h3>Location</h3>
      <div class="form-group col-sm-4 offset-4 center-block">
        <input v-model="zipcodeInput" type="text" id="zipcode-input" placeholder="zip code" class="form-control text-center">
      </div>
      <div class="form-group">
        <button class="btn btn-dark">Go!</button>
      </div>
    </form>
  </div>
</template>

<script>
export default {
  name: 'EventFinder',
  props: {
    msg: String
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
