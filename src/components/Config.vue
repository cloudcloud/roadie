<template>
  <v-container fluid grid-list-sm>
    <v-layout row wrap>

      <v-toolbar flat color="white">
        <v-toolbar-title>roadie</v-toolbar-title>
        <v-divider class="mx-2" inset vertical></v-divider>
        <span>config</span>
        <v-spacer />
      </v-toolbar>
    </v-layout>

    <v-layout row wrap>
      <!--v-container mt-0 pt-0 xs12>
        <v-flex mb-3>
          <v-data-table :headers="headers" :items="sources" class="elevation-1">
            <template slot="items" slot-scope="props">
              <td><router-link :to="props.item.href" :href="props.item.href">{{ props.item.name }}</router-link></td>
              <td>{{ props.item.location }}</td>
              <td>{{ props.item.type }}</td>

            </template>
          </v-data-table>
        </v-flex>
      </v-container-->

    </v-layout>
  </v-container>
</template>

<script>
//import { mapActions, mapMutations, mapGetters } from 'vuex';

//export default {
  //data: () => ({
    //headers: [
      //{text: 'Name', align: 'left', value: 'name'},
      //{text: 'Location', align: 'left', value: 'location'},
      //{text: 'Type', align: 'left', value: 'type'},
    //],
    //sources: [],
  //}),
  //props: [],
  //created() {
    //this.$store.dispatch('getSources').then(() => {
      //this.loadSources();
    //})
  //},
  //methods: {
    //loadSources() {
      //this.sources = this.$store.getters.allSources;
    //},
    //...mapMutations(['resetSources']),
    //...mapActions(['getSources']),
  //},
  //computed: {
    //...mapGetters(['allSources']),
  //},
//};
</script>

<style></style>
