<template>
  <v-container>
    <v-row>

      <v-col cols="12">
        <v-card shaped>

          <v-card-title>
            <v-row justify="center" align="center">
              <v-col cols="6" align="left">
                destinations
              </v-col>

              <v-col cols="6">
                <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details></v-text-field>
              </v-col>
            </v-row>
          </v-card-title>

          <v-data-table :headers="headers" :items="destinations" :search="search" class="elevation-1">
            <template v-slot:item.name="{ item }">
              <router-link :to="item.href">{{ item.name }}</router-link>
            </template>
          </v-data-table>

        </v-card>
      </v-col>

    </v-row>
  </v-container>
</template>

<script>
import { mapActions, mapMutations, mapGetters } from 'vuex';

export default {
  data: () => ({
    headers: [
      {text: 'Name', align: 'left', value: 'name'},
      {text: 'Location', align: 'left', value: 'config.location'},
      {text: 'Type', align: 'left', value: 'type'},
    ],
    destinations: [],
    search: '',
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
