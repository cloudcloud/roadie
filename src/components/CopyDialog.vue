<template>
  <v-btn small @click="copy" align="right">
    <v-icon :icon="`${mdiContentCopy}`"></v-icon> Copy
  </v-btn>

  <v-dialog v-model="dialog" max-width="500">
    <v-card :loading="loading" class="mx-auto">
      <v-card-title>Copy</v-card-title>
      <v-card-subtitle>Copying {{ entry_name }}. What is the desired Destination?</v-card-subtitle>
      <v-card-text>
        <v-combobox
          density="comfortable"
          outlined
          label="Destination"
          autofocus
          v-model="destination"
          item-title="name"
          item-value="name"
          :items="destinations">
        </v-combobox>
      </v-card-text>

      <v-card-actions>
        <v-spacer />
        <v-btn @click="close" color="error">Cancel</v-btn>
        <v-btn @click="save" color="primary">Copy</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import { mapGetters } from 'vuex';
import { mdiContentCopy } from '@mdi/js';

export default {
  data: () => ({
    mdiContentCopy,
    destinations: {},
    dialog: false,
    loading: false,
  }),
  props: ['source_name', 'entry_name'],
  created() {
    this.loadDestinations();
  },
  methods: {
    close() {
      this.dialog = false;
    },
    copy() {
      this.dialog = true;
    },
    created() {
      this.loadDestinations();
    },
    loadDestinations() {
      this.destinations = this.$store.getters.allDestinations;
    },
    save() {
      this.loading = true;
      this.$store.dispatch('pushCopy', {
        source_name: this.source_name,
        entry_name: this.entry_name,
        destination_name: this.destination.name,
      }).then(() => {
        this.loading = false;
        this.close();
      });
    },
  },
  computed: {
    ...mapGetters(['getCopyState']),
  },
};
</script>

<style></style>
