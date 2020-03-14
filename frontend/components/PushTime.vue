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
    <v-btn color="#007E33" class="mr-4" @click="modify">
      Submit
    </v-btn>
    <v-btn color="#B71C1C" class="mr-4" @click="reset">
      Reset
    </v-btn>
    <v-snackbar v-model="snackbar" color="#007E33" top>
      Pushtime has been updated!!!
      <v-btn color="white" text @click="snackbar = false">
        ×
      </v-btn>
    </v-snackbar>
  </v-form>
</template>

<script>
export default {
  data() {
    return {
      push_at: '',
      dialog: false,
      snackbar: false
    }
  },
  async mounted() {
    await this.$store.dispatch('push_time/fetch')
    this.push_at = this.$store.state.push_time.push_at
  },
  methods: {
    modify() {
      this.$store.dispatch('push_time/patch', this.push_at)
      this.snackbar = true
    },
    reset() {
      this.$store.commit('push_time/reset')
      this.push_at = ''
    }
  }
}
</script>
