import api from '../api/api';

class PasteService {
  async addPaste(title: string, content: string) {
    const { data } = await api.post('/new-paste', {
      title: title,
      content: content
    });

    return data;
  }
}

export const pasteService = new PasteService();
