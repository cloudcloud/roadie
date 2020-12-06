<template>
  <v-container fluid grid-list-sm>
    <v-layout row wrap>

      <v-toolbar flat color="white">
        <v-toolbar-title>roadie</v-toolbar-title>
        <v-divider class="mx-2" inset vertical></v-divider>
        <span>destination</span>
        <v-divider class="mx-2" inset vertical></v-divider>
        <span>{{ destination_name }}</span>
        <v-spacer />
      </v-toolbar>
    </v-layout>

    <v-layout row wrap>
      <v-container mt-0 pt-0 xs12>
        <v-flex mb-3>
          <v-data-table :headers="headers" :items="destination.entries" class="elevation-1">

            <template v-slot:item.action="{ item }">
              <v-btn block small @click="remove(item.entry)">Remove</v-btn>
            </template>

          </v-data-table>
        </v-flex>
      </v-container>

      <v-dialog v-model="dialog" max-width="500">
        <v-card :loading="loading" class="mx-auto">
          <v-card-title>
            Remove
          </v-card-title>

          <v-card-subtitle>
            Removing {{ entry }}. Are you sure?
          </v-card-subtitle>

          <v-card-text>
            ...
          </v-card-text>

          <v-card-actions>
            <v-card-actions>
              <v-spacer />
              <v-btn @click="close">No</v-btn>
              <v-btn @click="save" class="error">Yes</v-btn>
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
    destination: {},
    dialog: false,
    entry: '',
    loading: false,
  }),
  props: ['destination_name'],
  created() {
    this.loadDestination();
  },
  methods: {
    close() {
      this.dialog = false;
    },
    remove(entry) {
      this.entry = entry;
      this.dialog = true;
    },
    loadDestination() {
      this.$store.dispatch('getDestination', this.destination_name.trim().toLowerCase()).then(() => {
        this.destination = this.$store.getters.allDestination;
      });
    },
    save() {
      //this.loading = true;
      //var s = Object.assign({}, this.source.source);
      //s.entry = this.entry;

      //this.$store.dispatch('pushCopy', {
        //source: s,
        //destination: this.destination,
      //}).then(() => {
        //this.loading = false;
        //this.close();
      //});
    },
    ...mapMutations(['resetDestination']),
    ...mapActions(['getDestination']),
  },
  computed: {
    ...mapGetters(['allDestination']),
  },
};
</script>

<style></style>
