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
                <v-text-field v-model="search" single-line hide-details clearable placeholder="evil" variant="underlined">
                  <v-icon :icon="`${mdiMagnify}`"></v-icon>
                </v-text-field>
              </v-col>
            </v-row>
          </v-card-title>

          <v-data-table-virtual :headers="headers" :items="destinations" :search="search" hover class="elevation-1">
            <template v-slot:item.name="{ item }">
              <router-link :to="item.href">{{ item.name }}</router-link>
            </template>
          </v-data-table-virtual>

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
      {title: 'Name', align: 'left', key: 'name'},
      {title: 'Location', align: 'left', key: 'config.location'},
      {title: 'Type', align: 'left', key: 'type'},
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
