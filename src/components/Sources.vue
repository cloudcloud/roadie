<template>
  <v-container>
    <v-row>

      <v-col cols="12">
        <v-card shaped>

          <v-card-title>
            <v-row justify="center" align="center">
              <v-col cols="6" align="left">sources</v-col>

              <v-col cols="6">
                <v-text-field v-model="search" variant="underlined" clearable placeholder="evil" single-line hide-details>
                  <v-icon :icon="`${mdiMagnify}`"></v-icon>
                </v-text-field>
              </v-col>
            </v-row>
          </v-card-title>

          <v-data-table-virtual :headers="headers" :items="sources" :search="search" class="elevation-1" hover>
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
      {title: 'Bucket', align: 'left', key: 'config.bucket'},
      {title: 'Path', align: 'left', key: 'config.path'},
      {title: 'Type', align: 'left', key: 'type'},
    ],
    sources: [],
    search: '',
  }),
  props: [],
  created() {
    this.$store.dispatch('getSources').then(() => {
      this.loadSources();
    })
  },
  methods: {
    loadSources() {
      this.sources = this.$store.getters.allSources;
    },
    ...mapMutations(['resetSources']),
    ...mapActions(['getSources']),
  },
  computed: {
    ...mapGetters(['allSources']),
  },
};
</script>

<style></style>
