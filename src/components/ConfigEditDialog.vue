<template>
  <v-btn @click="edit" small>
    <v-icon :icon="`${mdiWrenchOutline}`"></v-icon> Edit
  </v-btn>

  <v-dialog v-model="dialog" max-width="500">
    <v-card class="mx-auto">
      <v-card-title>Edit</v-card-title>
      <v-card-subtitle>Editing "{{ name }}".</v-card-subtitle>
      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn @click="close" color="primary">Cancel</v-btn>
        <v-btn @click="run" color="error">Save</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import { mdiWrenchOutline } from '@mdi/js';

export default {
  data: () => ({
    mdiWrenchOutline,
    dialog: false,
  }),
  props: ['name', 'fullObj', 'type'],
  methods: {
    close() {
      this.dialog = false;
    },
    edit() {
      this.dialog = true;
    },
    run() {
    },
  },
};
</script>

<style></style>
