<template>
  <v-container>
    <v-row>

      <v-col cols="12">
        <v-card shaped>

          <v-card-title>
            <v-row justify="center" align="center">
              <v-col cols="6" align="left">
                <span>destination <v-icon :icon="`${mdiArrowRightThinCircleOutline}`"></v-icon> {{destination_name}}</span>
              </v-col>

              <v-col cols="6">
                <v-text-field v-model="search" single-line hide-details clearable variant="underlined">
                  <v-icon :icon="`${mdiMagnify}`"></v-icon>
                </v-text-field>
              </v-col>
            </v-row>
          </v-card-title>

          <v-data-table :headers="headers" :items="destination.entries" :search="search" class="elevation-1">
            <template v-slot:item.action="{ item }">
              <v-btn block small @click="remove(item.entry)">Remove</v-btn>
            </template>
          </v-data-table>

          <v-dialog v-model="dialog" max-width="500">
            <v-card :loading="loading" class="mx-auto">
              <v-card-title>
                Remove
              </v-card-title>

              <v-card-subtitle>
                Removing {{ entry }}. Are you sure?
              </v-card-subtitle>

              <v-card-actions>
                <v-card-actions>
                  <v-spacer />
                  <v-btn @click="close" color="primary">No</v-btn>
                  <v-btn @click="save" color="error">Yes</v-btn>
                </v-card-actions>
              </v-card-actions>

            </v-card>
          </v-dialog>

        </v-card>
      </v-col>

    </v-row>
  </v-container>
</template>

<script>
import { mapActions, mapMutations, mapGetters } from 'vuex';
import { mdiArrowRightThinCircleOutline, mdiMagnify } from '@mdi/js';

export default {
  data: () => ({
    mdiMagnify,
    mdiArrowRightThinCircleOutline,
    headers: [
      {title: 'Name', align: 'left', key: 'entry'},
      {title: 'Actions', key: 'action'},
    ],
    destination: {},
    dialog: false,
    entry: '',
    loading: false,
    search: '',
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
      this.loading = true;
      this.$store.dispatch('removeFile', {
        destination_name: this.destination_name,
        entry_name: this.entry,
      }).then(() => {
        this.loading = false;
        this.close();
        this.loadDestination();
      });
    },
    ...mapMutations(['resetDestination']),
    ...mapActions(['getDestination', 'removeFile']),
  },
  computed: {
    ...mapGetters(['allDestination']),
  },
};
</script>

<style></style>
