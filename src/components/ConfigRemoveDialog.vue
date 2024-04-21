<template>
  <v-btn @click="remove" color="error" small>
    <v-icon :icon="`${mdiTrashCanOutline}`"></v-icon> Delete
  </v-btn>

  <v-dialog v-model="dialog" max-width="500">
    <v-card class="mx-auto">
      <v-card-title>Remove</v-card-title>
      <v-card-subtitle>Removing "{{ name }}". Are you sure?</v-card-subtitle>
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
    },
  },
};
</script>

<style></style>
