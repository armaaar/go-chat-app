<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router';
import router from './router';
import { store } from './store';

function logout() {
  store.name = '';
  router.push('/');
}
</script>

<template>
  <div>
    <header class="header">
      <h1 class="header__title">Go Chat App</h1>
      <nav class="header__nav">
        <RouterLink class="header__nav-item" to="/">Home</RouterLink>
        <template v-if="store.isNameValid">
          <RouterLink class="header__nav-item" to="/chat">Chat</RouterLink>
          <a class="header__nav-item" @click.prevent="logout">Logout</a>
        </template>
      </nav>
    </header>

    <main class="page">
      <RouterView />
    </main>
  </div>
</template>

<style lang="scss" scoped>
.header {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  padding: var(--page-padding);

  border-bottom: 2px solid var(--color-border);
  margin-bottom: var(--section-gap);

  &__title {
    font-weight: bold;
  }

  &__nav {
    &-item {
      cursor: pointer;
      padding: 8px 16px;
      margin-right: 8px;
      text-decoration: none;

      &:last-child() {
        margin-right: 0;
      }
      &:hover,
      &:focus {
        border-bottom: 1px solid var(--color-active-selected);
      }
    }
  }
}

.page {
  padding: 0 var(--page-padding) var(--section-gap);
}
</style>
