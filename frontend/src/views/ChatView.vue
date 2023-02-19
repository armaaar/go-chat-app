<script setup lang="ts">
import axios from 'axios';
import router from '@/router';
import { store } from '@/store';
import { onMounted, onUnmounted, onUpdated, ref } from 'vue';
import stc from 'string-to-color';

if (!store.isNameValid) {
  router.push('/');
}

interface Message {
  ID: number;
  Name: string;
  Message: string;
  CreatedAt: string;
}

// Constants
const userName = store.name;
const maxMessageLength = 500;

// Messages element controls
const messagesEl = ref<null | HTMLDivElement>(null);
const stickMessagesBottom = ref(true);

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

// Messages controls
const wsConnection = ref<WebSocket | null>(null);
const messages = ref<Message[]>([]);
const newMessage = ref('');
const error = ref('');
const isDisabled = ref(true);

onMounted(() => {
  // Messages
  axios
    .get<Message[]>(`${import.meta.env.VITE_API_ROOT}/messages`)
    .then((response) => {
      messages.value = response.data.sort((a, b) => a.ID - b.ID);
    });

  // WebSockets
  wsConnection.value = new WebSocket(
    `${import.meta.env.VITE_WS_ROOT}/messages/ws`
  );

  wsConnection.value.onopen = () => {
    isDisabled.value = false;
    newMessage.value = '';
    error.value = '';
  };
  wsConnection.value.onerror = (ev) => {
    isDisabled.value = true;
    newMessage.value = '';
    error.value = 'An error occurred! Please refresh the page';
    console.error('WS Error:', ev);
  };
  wsConnection.value.onclose = () => {
    isDisabled.value = true;
    newMessage.value = '';
    error.value = 'Connection to the server is closed!';
  };
  wsConnection.value.onmessage = (ev) => {
    const msg = JSON.parse(ev.data) as Message;
    messages.value.push(msg);
  };
});
onUnmounted(() => {
  if (wsConnection.value) {
    wsConnection.value.close();
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

  if (wsConnection.value) {
    wsConnection.value.send(
      JSON.stringify({ Name: userName, Message: newMessage.value })
    );
  }
  error.value = '';
  newMessage.value = '';
  scrollMessagesToBottom();
}
</script>

<template>
  <h2 class="title">
    Start Chatting, <span :style="{ color: stc(userName) }">{{ userName }}</span
    >!
  </h2>
  <div class="messages" ref="messagesEl">
    <div
      v-for="message in messages"
      :key="message.ID"
      class="messages__item"
      :class="{ 'messages__item--active': message.Name === userName }"
    >
      <span class="messages__item-time"
        >[{{ new Date(message.CreatedAt).toLocaleString() }}]</span
      >
      &nbsp;
      <span class="messages__item-name" :style="{ color: stc(message.Name) }"
        >{{ message.Name }}:</span
      >
      &nbsp;
      <span class="messages__item-message">{{ message.Message }}</span>
    </div>
  </div>
  <div class="new-message">
    <input
      class="new-message__field"
      type="text"
      v-model="newMessage"
      :maxlength="maxMessageLength"
      @keypress.enter="sendMessage"
      :disabled="isDisabled"
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
