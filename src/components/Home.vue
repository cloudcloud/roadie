<template>
  <v-container>
    <v-row>

      <v-col cols="12">
        <v-card shaped>

          <v-card-title>
            <v-row justify="center" align="center">
              <v-col cols="6" align="left">roadie</v-col>

              <v-col cols="6">
                <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details />
              </v-col>
            </v-row>
          </v-card-title>

          <v-data-table :headers="headers" :items="historical" :search="search" class="elevation-1">
            <template v-slot:item.source.name="{ item }">
              <router-link :to="item.source.href">{{ item.source.name }}</router-link>
            </template>
            <template v-slot:item.destination.name="{ item }">
              <router-link :to="item.destination.href">{{ item.destination.name }}</router-link>
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
      {text: 'Source', align: 'left', value: 'source.name'},
      {text: 'Destination', align: 'left', value: 'destination.name'},
      {text: 'Pattern', align: 'left', value: 'pattern'},
      {text: 'Date', align: 'left', value: 'occurred_at'},
    ],
    historical: [],
    search: '',
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
