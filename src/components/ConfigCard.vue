<template>
  <v-card shaped :title="title" :subtitle="subtitle">
    <template v-slot:append>
      <ConfigAddDialog :addType="type"></ConfigAddDialog>
    </template>

    <v-card-text>
      <v-data-table-virtual :items="items" :items-per-page="0" :headers="headers" hover no-filter disable-pagination>
        <template v-slot:item.action="{ item }">
          <v-spacer></v-spacer>
          <ConfigEditDialog :name="item.name" :fullObj="item" :type="type"></ConfigEditDialog>
          <ConfigRemoveDialog :name="item.name" :fullObj="item" :type="type"></ConfigRemoveDialog>
          <v-spacer></v-spacer>
        </template>

        <template v-slot:item.name="{ item }">
          <router-link :to="makeHref(item.name)">{{ item.name }}</router-link>
        </template>

        <template v-slot:item.disk_free="{ item }">
          <strong v-if="item.disk_free">{{ item.disk_free }}</strong>
        </template>

        <template #bottom></template>
      </v-data-table-virtual>
    </v-card-text>
  </v-card>
</template>

<script>
import ConfigAddDialog from './ConfigAddDialog';
import ConfigEditDialog from './ConfigEditDialog';
import ConfigRemoveDialog from './ConfigRemoveDialog';
import { mdiPlusOutline } from '@mdi/js';

export default {
  data: () => ({
    mdiPlusOutline,
    headers: [
      {title: 'Name', align: 'left', key: 'name'},
      {title: 'Location', align: 'left', key: 'config.location'},
      {title: 'Disk Free', align: 'left', key: 'disk_free'},
      {title: 'Actions', align: 'center', key: 'action', sortable: false},
    ],
  }),
  props: ['title', 'subtitle', 'items', 'link-prefix', 'type'],
  methods: {
    makeHref(val) {
      return this.linkPrefix + val;
    },
  },
  components: {
    ConfigAddDialog,
    ConfigEditDialog,
    ConfigRemoveDialog,
  },
};
</script>

<style></style>

