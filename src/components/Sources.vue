<template>
  <v-container>
    <v-row>

      <v-col cols="12">
        <v-card shaped>

          <v-card-title>
            <v-row justify="center" align="center">
              <v-col cols="6" align="left">sources</v-col>

              <v-col cols="6">
                <v-text-field v-model="search" append-icon="mdi-magnify" label="Search" single-line hide-details />
              </v-col>
            </v-row>
          </v-card-title>

          <v-data-table :headers="headers" :items="sources" :search="search" class="elevation-1">
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
      {text: 'Bucket', align: 'left', value: 'config.bucket'},
      {text: 'Path', align: 'left', value: 'config.path'},
      {text: 'Type', align: 'left', value: 'type'},
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
