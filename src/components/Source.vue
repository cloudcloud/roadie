<template>
  <v-container fluid grid-list-sm>
    <v-layout row wrap>

      <v-toolbar flat color="white">
        <v-toolbar-title>roadie</v-toolbar-title>
        <v-divider class="mx-2" inset vertical></v-divider>
        <span>source</span>
        <v-divider class="mx-2" inset vertical></v-divider>
        <span>{{ source_name }}</span>
        <v-spacer />
      </v-toolbar>
    </v-layout>

    <v-layout row wrap>
      <v-container mt-0 pt-0 xs12>
        <v-flex mb-3>
          <v-data-table :headers="headers" :items="source.entries" class="elevation-1">

            <template v-slot:item.action="{ item }">
              <v-btn block small @click="copy(item.entry)">Copy</v-btn>
            </template>

          </v-data-table>
        </v-flex>
      </v-container>

      <v-dialog v-model="dialog" max-width="500">
        <v-card :loading="loading" class="mx-auto">
          <v-card-title>
            Copy
          </v-card-title>

          <v-card-subtitle>
            Copying {{ entry }}. What is the desired Destination?
          </v-card-subtitle>

          <v-card-text>
            <v-combobox dense outlined persistent-hint solo autofocus item-text="location" :items="destinations"></v-combobox>
          </v-card-text>

          <v-card-actions>
            <v-card-actions>
              <v-spacer />
              <v-btn @click="close">Cancel</v-btn>
              <v-btn @click="save">Save</v-btn>
            </v-card-actions>
          </v-card-actions>

        </v-card>
      </v-dialog>

    </v-layout>
  </v-container>
</template>

<script>
import { mapActions, mapMutations, mapGetters } from 'vuex';

export default {
  data: () => ({
    headers: [
      {text: 'Name', align: 'left', value: 'entry'},
      {text: 'Actions', value: 'action'},
    ],
    source: {},
    dialog: false,
    entry: '',
    loading: false,
    destinations: [],
  }),
  props: ['source_name'],
  created() {
    this.$store.dispatch('getSource', this.source_name.trim().toLowerCase()).then(() => {
      this.loadSource();
    })
  },
  methods: {
    close() {
      this.dialog = false;
    },
    copy(entry) {
      this.entry = entry;
      this.dialog = true;
      this.destinations = this.$store.getters.allDestinations;
    },
    loadSource() {
      this.source = this.$store.getters.allSource;
    },
    save() {
      this.loading = true;
      this.close();
      this.loading = false;
    },
    ...mapMutations(['resetSource']),
    ...mapActions(['getSource']),
  },
  computed: {
    ...mapGetters(['allSource']),
  },
};
</script>

<style></style>
