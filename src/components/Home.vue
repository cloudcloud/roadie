<template>
  <v-container>
    <v-row>

      <v-col cols="12">
        <v-card shaped>

          <v-card-title>
            <v-row justify="center" align="center">
              <v-col cols="6" align="left">roadie</v-col>

              <v-col cols="6">
                <v-text-field v-model="search" clearable variant="underlined" single-line hide-details>
                  <v-icon :icon="`${mdiMagnify}`"></v-icon>
                </v-text-field>
              </v-col>
            </v-row>
          </v-card-title>

          <v-data-table :headers="headers" :items="historical" :search="search" class="elevation-1" hover height="100%" items-per-page="20">
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
import { mdiMagnify } from '@mdi/js';

export default {
  data: () => ({
    mdiMagnify,
    headers: [
      {title: 'Source', align: 'left', key: 'source.name'},
      {title: 'Destination', align: 'left', key: 'destination.name'},
      {title: 'Pattern', align: 'left', key: 'pattern'},
      {title: 'Date', align: 'left', key: 'occurred_at'},
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
