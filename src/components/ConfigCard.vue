<template>
  <v-card shaped>
    <v-card-title><span>{{ title }}</span></v-card-title>
    <v-card-subtitle>{{ subtitle }}</v-card-subtitle>

    <v-card-text>
      <v-data-table-virtual :items="items" :items-per-page="0" :headers="headers" hover no-filter disable-pagination>
        <template v-slot:item.action="{ item }">
          <v-spacer></v-spacer>
          <v-btn @click="edit">
            <v-icon :icon="`${mdiWrenchOutline}`"></v-icon> Edit
          </v-btn>
          <v-btn @click="remove" color="error">
            <v-icon :icon="`${mdiTrashCanOutline}`"></v-icon> Delete
          </v-btn>
          <v-spacer></v-spacer>
        </template>

        <template v-slot:item.name="{ item }">
          <router-link :to="makeHref(item.name)">{{ item.name }}</router-link>
        </template>

        <template v-slot:item.disk_info="{ item }">
          <strong v-if="item.disk_info">{{ diskSize(item.disk_info.free) }}</strong>
        </template>

        <template #bottom></template>
      </v-data-table-virtual>
    </v-card-text>
  </v-card>

  <v-dialog v-model="editDialog" max-width="500">
  </v-dialog>

  <v-dialog v-model="removeDialog" max-width="500">
  </v-dialog>
</template>

<script>
import { mdiWrenchOutline, mdiTrashCanOutline } from '@mdi/js';

export default {
  data: () => ({
    mdiWrenchOutline,
    mdiTrashCanOutline,
    editDialog: false,
    removeDialog: false,
  }),
  props: ['title', 'subtitle', 'items', 'headers', 'link-prefix'],
  methods: {
    diskSize(num) {
      let n = num
      if (n / 1024 > 1) {
        n /= 1024
        if (n / 1024 > 1) {
          n /= 1024
          if (n / 1024 > 1) {
            n /= 1024
            if (n / 1024 > 1) {
              n /= 1024
              if (n / 1024 > 1) {
                n /= 1024
                return n.toLocaleString('en-AU', { style: 'unit', unit: 'petabyte'});
              }
              return n.toLocaleString('en-AU', { style: 'unit', unit: 'terabyte'});
            }
            return n.toLocaleString('en-AU', { style: 'unit', unit: 'gigabyte' });
          }
          return n.toLocaleString('en-AU', { style: 'unit', unit: 'megabyte' });
        }
        return n.toLocaleString('en-AU', { style: 'unit', unit: 'kilobyte' });
      }
      return n.toLocaleString('en-AU', { style: 'unit', unit: 'byte' });
    },
    makeHref(val) {
      return this.linkPrefix + val;
    },

    close() {
      this.editDialog = false;
      this.removeDialog = false;
    },
    edit(name) {
      // type cannot be changed, remove and create new
      this.editDialog = true;
    },
    remove(name) {
      this.removeDialog = true;
    },
  },
};
</script>

<style></style>

