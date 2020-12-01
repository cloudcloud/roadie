<template>
  <v-container fluid grid-list-sm>
    <v-layout row wrap>

      <v-toolbar flat color="white">
        <v-toolbar-title>roadie</v-toolbar-title>
        <v-divider class="mx-2" inset vertical></v-divider>
        <span>destinations</span>
        <v-spacer />
      </v-toolbar>
    </v-layout>

    <v-layout row wrap>
      <v-container mt-0 pt-0 xs12>
        <v-flex mb-3>
          <v-data-table :headers="headers" :items="destinations" class="elevation-1">

            <template v-slot:item.name="{ item }">
              <router-link :to="item.href">{{ item.name }}</router-link>
            </template>

          </v-data-table>
        </v-flex>
      </v-container>

    </v-layout>
  </v-container>
</template>

<script>
import { mapActions, mapMutations, mapGetters } from 'vuex';

export default {
  data: () => ({
    headers: [
      {text: 'Name', align: 'left', value: 'name'},
      {text: 'Location', align: 'left', value: 'location'},
      {text: 'Type', align: 'left', value: 'type'},
    ],
    destinations: [],
  }),
  props: [],
  created() {
    this.$store.dispatch('getDestinations').then(() => {
      this.loadDestinations();
    })
  },
  methods: {
    loadDestinations() {
      this.destinations = this.$store.getters.allDestinations;
    },
    ...mapMutations(['resetDestinations']),
    ...mapActions(['getDestinations']),
  },
  computed: {
    ...mapGetters(['allDestinations']),
  },
};
</script>

<style></style>
