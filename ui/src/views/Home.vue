<template>
  <v-container fluid>
    <v-slide-y-transition mode="out-in">
      <v-layout column align-center>
        <img src="@/assets/gh-ost-logo-light-160.png" alt="gh-ost">
        <v-subheader>Migrations in Progress</v-subheader>
        <MigrationsInProgress/>
        <v-subheader>Pending Migrations</v-subheader>
      </v-layout>
    </v-slide-y-transition>
  </v-container>
</template>

<script>
import MigrationsInProgress from '@/components/MigrationsInProgress.vue';

export default {
  name: 'Home',
  components: {
    MigrationsInProgress
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1,
h2 {
  font-weight: normal;
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
