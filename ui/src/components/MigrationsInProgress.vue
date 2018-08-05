<template>
  <v-data-table
    :items="migrations"
  >
    <template slot="headers" slot-scope="props">
      <tr>
        <th>id</th>
        <th>started at</th>
        <th>cluster name</th>
        <th>requestor</th>
        <th>alter statement</th>
        <th>approved by</th>
        <th>copy progress</th>
        <th>status</th>
      </tr>
    </template>
    <template slot="items" slot-scope="props">
    </template>
  </v-data-table>
</template>

<script>
export default {
  name: 'MigartionsInProgress',
  props: {
    migrations: []
  }
};
</script>
