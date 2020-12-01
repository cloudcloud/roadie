<template>
  <v-container fluid grid-list-sm>
    <v-layout row wrap>
      <v-toolbar flat color="white">
        <v-toolbar-title>roadie</v-toolbar-title>
        <v-divider class="mx-2" inset vertical></v-divider>
        <span class="mr-2">history</span>
        <v-spacer />
      </v-toolbar>
    </v-layout>

    <v-layout row wrap>
      <v-container mt-0 pt-0 xs12>
        <v-flex mb-3>
          <v-data-table :headers="headers" :items="historical" class="elevation-1">

            <template v-slot:item.source.location="{ item }">
              <router-link :to="item.source.href">{{ item.source.name }}</router-link>
            </template>

            <template v-slot:item.destination.location="{ item }">
              <router-link :to="item.destination.href">{{ item.destination.name }}</router-link>
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
      {text: 'Source', align: 'left', value: 'source.location'},
      {text: 'Destination', align: 'left', value: 'destination.location'},
      {text: 'Pattern', align: 'left', value: 'pattern'},
      {text: 'Date', align: 'left', value: 'date'},
    ],
    historical: [],
  }),
  props: [],
  created() {
    this.$store.dispatch('getHistorical').then(() => {
      this.loadHistorical();
    })
  },
  methods: {
    loadHistorical() {
      this.historical = this.$store.getters.allHistorical;
    },
    ...mapMutations(['resetHistorical']),
    ...mapActions(['getHistorical']),
  },
  computed: {
    ...mapGetters(['allHistorical']),
  },
};
</script>

<style></style>
