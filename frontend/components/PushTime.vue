<template>
  <v-form ref="form">
    <v-text-field
      v-model="push_at"
      :counter="10"
      label="Push Time"
      required
    ></v-text-field>

    <v-btn color="success" class="mr-4" @click="modify">
      Submit
    </v-btn>

    <v-btn color="error" class="mr-4" @click="reset">
      Reset
    </v-btn>
  </v-form>
</template>

<script>
// import { mapState } from 'vuex'

export default {
  data() {
    return {
      // Inputに入力するときは変更せずに、ボタンを押したときだけ変更にしたい！そしてStateは自動的に反映されてほしい。
      push_at: this.$store.state.push_time.push_at
    }
  },
  //   computed: {
  //     ...mapState({
  //       push_at: (state) => state.push_time.push_at
  //     })
  //   },
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
