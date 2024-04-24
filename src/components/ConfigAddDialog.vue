<template>
  <v-btn @click="add" variant="text" small color="primary">
    <v-icon :icon="`${mdiPlusOutline}`"></v-icon> Add
  </v-btn>

  <v-dialog v-model="dialog" max-width="500">
    <v-card
      title="Add"
      :subtitle="'Add new ' + addType + '.'"
      class="mx-auto"
      width="500">

      <v-card-text>
        <v-form ref="add">
          <v-text-field
            v-model="name"
            :rules="[rules.required, rules.name]"
            label="Name"
            required />

          <v-select
            v-model="select"
            :items="selectOptions"
            :rules="[rules.required]"
            label="Item"
            required />

          <div v-if="select === 'local_path'">
            <v-text-field
              v-model="path"
              :rules="[rules.required]"
              label="Path" />
          </div>
          <div v-else-if="select === 's3'">
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
              label="Depth"
              placeholder="0" />
          </div>

        </v-form>
      </v-card-text>

      <v-card-actions>
        <v-spacer></v-spacer>
        <v-btn @click="close" color="primary">Cancel</v-btn>
        <v-btn @click="run" color="error">Save</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import { mapActions } from 'vuex';
import { mdiPlusOutline } from '@mdi/js';

export default {
  data: () => ({
    mdiPlusOutline,
    dialog: false,
    name: '',
    rules: {
      required: v => !!v || 'Required.',
      name: v => !/[^A-Za-z0-9\-_]+/.test(v) || 'Name should only contain letters, numbers, "-", and "_".',
    },
    select: 'local_path',
    selectOptions: [
      'local_path',
      's3',
    ],
    path: '',
    bucket: '',
    depth: 0,
  }),
  props: ['addType'],
  beforeUpdate() {
    if (this.addType === 'source') {
      this.selectOptions = ['local_path', 's3'];
    } else if (this.addType === 'destination') {
      this.selectOptions = ['local_path'];
    }
  },
  methods: {
    add() {
      this.dialog = true;
    },
    close() {
      this.dialog = false;
    },
    run() {
      this.loading = true;
      let payload = {
        name: this.name,
        type: this.select,
        path: this.path,
      };
      if (this.select === 's3') {
        payload.bucket = this.bucket;
        payload.depth = parseInt(this.depth);
      }

      this.$store.dispatch('addConfig', {
        type: this.addType,
        payload: payload,
      }).then(() => {
        this.loading = false;
        this.close();
      });
    },
    ...mapActions(['addConfig']),
  },
};
</script>

<style></style>
