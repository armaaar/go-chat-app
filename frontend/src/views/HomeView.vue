<script setup lang="ts">
import router from '@/router';
import { store, getNameError, nameMaxLength } from '@/store';
import { ref } from 'vue';

const name = ref(store.name);
const error = ref('');

function registerName() {
  const nameError = getNameError(name.value);
  if (nameError) {
    error.value = nameError;
    return;
  }

  store.name = name.value;
  router.push('/chat');
}
</script>

<template>
  <section class="home">
    <h2 class="home__title">Welcome to Go Chat App!</h2>
    <p>Please enter your name:</p>
    <input
      class="home__name-field"
      type="text"
      v-model="name"
      :maxlength="nameMaxLength"
      @keypress.enter="registerName"
    />
    <p v-if="error" class="home__error">{{ error }}</p>
    <button class="home__submit-btn" @click="registerName">Chat!</button>
  </section>
</template>

<style lang="scss" scoped>
.home {
  text-align: center;

  & > * {
    display: block;
    margin: 0 auto 16px;
    text-align: center;
    &:last-child {
      margin-bottom: 0;
    }
  }

  &__title {
    font-weight: bold;
  }

  &__name-field {
    padding: 8px 16px;
  }

  &__error {
    color: var(--color-error);
  }

  &__submit-btn {
    padding: 8px 16px;
  }
}
</style>
