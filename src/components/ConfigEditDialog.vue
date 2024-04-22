<template>
  <v-btn @click="edit" small class="mx-1">
    <v-icon :icon="`${mdiWrenchOutline}`"></v-icon> Edit
  </v-btn>

  <v-dialog v-model="dialog" max-width="500">
    <v-card
      title="Edit"
      :subtitle="'Editing ' + name + '.'"
      class="mx-auto">

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
