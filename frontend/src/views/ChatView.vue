<script setup lang="ts">
import router from '@/router';
import { store } from '@/store';
import { onMounted, onUpdated, ref } from 'vue';
import stc from 'string-to-color';

if (!store.isNameValid) {
  router.push('/');
}

interface Message {
  id: number;
  name: string;
  message: string;
  created_at: string;
}

const userName = store.name;
const maxMessageLength = 500;

const messagesEl = ref<null | HTMLDivElement>(null);
const messages = ref<Message[]>([
  {
    id: 1,
    name: userName,
    message:
      'hahaha hahaha hahaha hahaha hahaha hahaha hahaha hahaha hahaha hahaha hahaha hahaha hahaha hahaha hahaha hahaha hahaha hahaha hahaha hahaha hahaha ',
    created_at: 'xxxx',
  },
  {
    id: 2,
    name: 'userName',
    message: 'hehehe',
    created_at: 'xxsx',
  },
  ...Array(300).fill({
    id: 3,
    name: 'asd',
    message: 'hohohoh',
    created_at: 'xxdx',
  }),
]);
const stickMessagesBottom = ref(true);
const newMessage = ref('');
const error = ref('');

function scrollMessagesToBottom() {
  if (messagesEl.value) {
    messagesEl.value.scrollTop = messagesEl.value.scrollHeight;
  }
}

onMounted(() => {
  if (messagesEl.value) {
    scrollMessagesToBottom();
    messagesEl.value.onscroll = () => {
      if (messagesEl.value) {
        stickMessagesBottom.value =
          Math.abs(
            messagesEl.value.scrollHeight -
              messagesEl.value.scrollTop -
              messagesEl.value.clientHeight
          ) < 1;
      }
    };
  }
});

onUpdated(() => {
  if (stickMessagesBottom.value) {
    scrollMessagesToBottom();
  }
});

function sendMessage() {
  if (!newMessage.value) {
    error.value = 'Please enter a message first!';
    return;
  }
  if (newMessage.value.length > maxMessageLength) {
    error.value = `The message can't be longer than ${maxMessageLength} characters!`;
    return;
  }

  messages.value = [
    ...messages.value,
    {
      id: messages.value.length,
      name: userName,
      message: newMessage.value,
      created_at: 'xxxx',
    },
  ];
  newMessage.value = '';
  scrollMessagesToBottom();
}
</script>

<template>
  <h2 class="title">Start Chatting!</h2>
  <div class="messages" ref="messagesEl">
    <div
      v-for="message in messages"
      :key="message.id"
      class="messages__item"
      :class="{ 'messages__item--active': message.name === userName }"
    >
      <span class="messages__item-time">[{{ message.created_at }}]</span>
      &nbsp;
      <span class="messages__item-name" :style="{ color: stc(message.name) }"
        >{{ message.name }}:</span
      >
      &nbsp;
      <span class="messages__item-message">{{ message.message }}</span>
    </div>
  </div>
  <div class="new-message">
    <input
      class="new-message__field"
      type="text"
      v-model="newMessage"
      :maxlength="maxMessageLength"
      @keypress.enter="sendMessage"
    />
    <button class="new-message__btn" @click="sendMessage">Send</button>
    <p v-if="error" class="new-message__error">{{ error }}</p>
  </div>
</template>

<style lang="scss" scoped>
.title {
  text-align: center;
  margin-bottom: var(--section-gap);
}

.messages {
  border: 2px solid var(--color-border);
  background: var(--color-background-soft);
  min-height: 300px;
  height: 700px;
  max-height: 60vh;
  overflow-y: auto;

  &__item {
    margin-bottom: 8px;
    &--active {
      font-weight: bold;
    }
  }
}

.new-message {
  &__field {
    padding: 8px 16px;
    width: calc(100% - 100px);

    border: 2px solid var(--color-border);
    border-top-right-radius: 0;
    border-bottom-right-radius: 0;
  }

  &__btn {
    padding: 8px 16px;
    width: 100px;

    border: 2px solid var(--color-border);
    border-top-left-radius: 0;
    border-bottom-left-radius: 0;
  }

  &__error {
    display: block;
    width: 100%;
    color: var(--color-error);
  }
}
</style>
