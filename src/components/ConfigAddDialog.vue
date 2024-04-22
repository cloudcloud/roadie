<template>
  <v-btn @click="add" variant="text" small color="primary">
    <v-icon :icon="`${mdiPlusOutline}`"></v-icon> Add
  </v-btn>

  <v-dialog v-model="dialog" max-width="500">
    <v-card
      title="Add"
      :subtitle="'Add new ' + type + '.'"
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
import { mdiPlusOutline } from '@mdi/js';

export default {
  data: () => ({
    mdiPlusOutline,
    dialog: false,
  }),
  props: ['type'],
  methods: {
    add() {
      this.dialog = true;
    },
    close() {
      this.dialog = false;
    },
    run() {
    },
  },
};
</script>

<style></style>
