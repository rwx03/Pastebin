import api from '../api/api';
import { IPaste } from '../models/PasteModel';

class PasteService {
  async addPaste(title: string, content: string) {
    try {
      const { data } = await api.post('/paste', {
        title: title,
        content: content
      });

      return data;
    } catch (e) {
      console.log(e);
      throw e;
    }
  }

  async getPaste(id: string): Promise<IPaste> {
    try {
      const res = await api.get<IPaste>(`/paste?id=${id}`);
      return res.data;
    } catch (e) {
      console.log(e);
      throw e;
    }
  }
}

export const pasteService = new PasteService();
