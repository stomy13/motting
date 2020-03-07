<template>
  <v-form ref="form">
    <v-dialog v-model="dialog" width="290px">
      <template v-slot:activator="{ on }">
        <v-text-field
          v-model="push_at"
          label="Push Time"
          required
          readonly
          v-on="on"
        ></v-text-field>
      </template>
      <v-time-picker
        v-model="push_at"
        class="mt-2"
        format="24hr"
      ></v-time-picker>
    </v-dialog>
    <v-btn color="success" class="mr-4" @click="modify">
      Submit
    </v-btn>
    <v-btn color="error" class="mr-4" @click="reset">
      Reset
    </v-btn>
  </v-form>
</template>

<script>
export default {
  data() {
    return {
      push_at: this.$store.state.push_time.push_at,
      dialog: false
    }
  },
  methods: {
    modify() {
      this.$store.commit('push_time/modify', this.push_at)
    },
    reset() {
      this.$store.commit('push_time/reset')
      this.push_at = ''
    }
  }
}
</script>
