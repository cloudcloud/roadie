<template>
  <v-data-table :items="sub_source.entries" :items-per-page="0" :headers="headers" hover no-filter disable-pagination hide-default-header>
    <template v-slot:item.action="{ item }">
      <CopyDialog :source_name="source_name" :entry_name="sub_name + '/' + item.entry + '/'" />
    </template>
    <template v-slot:headers="{ }"><!-- skip! --></template>
    <template #bottom></template>
  </v-data-table>
</template>

<script>
import { mapActions, mapMutations, mapGetters } from 'vuex';
import { mdiMagnify } from '@mdi/js';
import CopyDialog from './CopyDialog';

export default {
  data: () => ({
    mdiMagnify,
    headers: [
      {sortable: false, key: 'entry'},
      {sortable: false, key: 'action'},
    ],
    sub_source: {},
    entry: '',
  }),
  props: ['source_name', 'sub_name'],
  created() {
    this.loadSource();
  },
  methods: {
    loadSource() {
      let data = {};
      data.source_name = this.source_name;
      data.sub_name = this.sub_name;
      this.$store.dispatch('getSubSource', data).then(() => {
        this.sub_source = this.$store.getters.allSubSource;
      });
    },
    ...mapMutations(['resetSubSource']),
    ...mapActions(['getSubSource']),
  },
  computed: {
    ...mapGetters(['allSubSource']),
  },
  components: {
    CopyDialog,
  },
};
</script>

<style></style>
