<template>
  <v-btn @click="remove" color="error" small class="mx-1">
    <v-icon :icon="`${mdiTrashCanOutline}`"></v-icon> Delete
  </v-btn>

  <v-dialog v-model="dialog" max-width="500">
    <v-card
      title="Remove"
      :subtitle="'Removing ' + name + '. Are you sure?'"
      :loading="loading"
      class="mx-auto">

      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn @click="close" color="primary">No</v-btn>
        <v-btn @click="run" color="error">Yes</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import { mdiTrashCanOutline } from '@mdi/js';

export default {
  data: () => ({
    mdiTrashCanOutline,
    dialog: false,
    loading: false,
  }),
  props: ['name', 'fullObj', 'type'],
  methods: {
    close() {
      this.dialog = false;
    },
    remove() {
      this.dialog = true;
    },
    run() {
      this.loading = true;
      this.$store.dispatch('removeFromConfig', type, name).then(() => {
        this.loading = false;
        this.close();
      });
    },
  },
};
</script>

<style></style>
