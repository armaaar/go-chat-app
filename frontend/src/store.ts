import { reactive } from 'vue';

export const nameMaxLength = 20;

export function getNameError(name: string): string | null {
  if (!name) {
    return 'Please enter a name!';
  }
  if (name.length > nameMaxLength) {
    return `name can't be longer than ${nameMaxLength} characters!`;
  }

  return null;
}

export const store = reactive({
  get name(): string {
    return localStorage.getItem('name') ?? '';
  },
  set name(newName: string) {
    localStorage.setItem('name', newName);
  },
  get isNameValid() {
    return !getNameError(this.name);
  },
});
