<template>
  <v-btn @click="edit" small class="mx-1">
    <v-icon :icon="`${mdiWrenchOutline}`"></v-icon> Edit
  </v-btn>

  <v-dialog v-model="dialog" max-width="500">
    <v-card
      title="Edit"
      :subtitle="'Editing ' + name + '.'"
      width="500"
      class="mx-auto">

      <v-form ref="edit">
        <v-text-field
          :model-value="name"
          disabled
          label="Name" />

        <v-text-field
          :model-value="fullObj.type"
          disabled
          label="Type" />

        <div v-if="fullObj.type === 'local_path'">
          <v-text-field
            v-model="path"
            :rules="[rules.required]"
            label="Path" />
        </div>
        <div v-else-if="fullObj.type === 's3'">
          <v-text-field
            v-model="path"
            :rules="[rules.required]"
            label="Path" />

          <v-text-field
            v-model="bucket"
            :rules="[rules.required, rules.name]"
            label="Bucket" />

          <v-text-field
            v-model="depth"
            :rules="[rules.required]"
            label="Depth" />
        </div>
      </v-form>

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
    path: '',
    bucket: '',
    depth: 0,
    rules: {
      required: v => !!v || 'Required.',
      name: v => !/[^A-Za-z0-9\-_]+/.test(v) || 'Name should only contain letters, numbers, "-", and "_".',
    },
  }),
  props: ['name', 'fullObj', 'type'],
  beforeUpdate() {
    if (this.fullObj.type === 'local_path') {
      this.path = this.fullObj.config.location;
    } else if (this.fullObj.type === 's3') {
      this.path = this.fullObj.config.path;
      this.bucket = this.fullObj.config.bucket;
      this.depth = this.fullObj.config.depth;
    }
  },
  methods: {
    close() {
      this.dialog = false;
    },
    edit() {
      this.dialog = true;
    },
    run() {
      this.loading = true;
      let payload = {
        name: this.name,
        path: this.path,
        bucket: this.bucket,
        depth: parseInt(this.depth),
      };
      this.$store.dispatch('editConfig', {
        type: this.type,
        name: payload.name,
        payload: payload,
      }).then(() => {
        this.loading = false;
        this.close();
      });
    },
  },
};
</script>

<style></style>
