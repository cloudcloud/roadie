<template>
  <v-data-table :items="sub_source.entries" :items-per-page="0" :headers="headers" hover no-filter disable-pagination hide-default-header>
    <template v-slot:item.action="{ item }">
      <v-btn block small @click="copy(item.entry)">
        <v-icon :icon="`${mdiMagnify}`"></v-icon> Copy
      </v-btn>
    </template>
    <template v-slot:headers="{ }"><!-- skip! --></template>
    <template #bottom></template>
  </v-data-table>

  <v-dialog v-model="dialog" max-width="500">
    <v-card :loading="loading" class="mx-auto">
      <v-card-title>Copy</v-card-title>
      <v-card-subtitle>Copying {{ entry }}. What is the desired Destination?</v-card-subtitle>
      <v-card-text>
        <v-combobox dense outlined persistent-hint solo autofocus v-model="destination" item-text="name" :items="destinations"></v-combobox>
      </v-card-text>

      <v-card-actions>
        <v-spacer />
        <v-btn @click="close">Cancel</v-btn>
        <v-btn @click="save">Copy</v-btn>
      </v-card-actions>
    </v-card>
  </v-dialog>
</template>

<script>
import { mapActions, mapMutations, mapGetters } from 'vuex';
import { mdiMagnify } from '@mdi/js';

export default {
  data: () => ({
    mdiMagnify,
    headers: [
      {sortable: false, key: 'entry'},
      {sortable: false, key: 'action'},
    ],
    sub_source: {},
    dialog: false,
    entry: '',
    loading: false,
    destinations: [],
    destination: '',
  }),
  props: ['source_name', 'sub_name'],
  created() {
    this.loadSource();
  },
  methods: {
    close() {
      this.dialog = false;
    },
    copy(entry) {
      this.entry = entry;
      this.dialog = true;
      this.destinations = this.$store.getters.allDestinations;
    },
    loadSource() {
      let data = {};
      data.source_name = this.source_name;
      data.sub_name = this.sub_name;
      this.$store.dispatch('getSubSource', data).then(() => {
        this.sub_source = this.$store.getters.allSubSource;
      });
    },
    save() {
      this.loading = true;
      this.$store.dispatch('pushCopy', {
        source_name: this.sub_source.source.name,
        entry_name: this.sub_name + '/' + this.entry,
        destination_name: this.destination.name,
      }).then(() => {
        this.loading = false;
        this.close();
      });
    },
    ...mapMutations(['resetSubSource']),
    ...mapActions(['getSubSource']),
  },
  computed: {
    ...mapGetters(['allSubSource', 'getCopyState']),
  },
};
</script>

<style></style>
