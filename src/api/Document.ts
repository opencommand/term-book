import axios from 'axios';
import config from './config';

const apiClient = axios.create({
  baseURL: config.API_BASE_URL,
  headers: { 'Content-Type': 'application/json' }
});

// get fill list
export const getFileListApi = async () => {
  try {
    const { data } = await apiClient.get('/document/list');

    if (!Array.isArray(data.data)) {
      throw new Error('返回数据格式错误');
    }

    return {
      success: true,
      data: data.data,
      message: data.msg || '获取成功'
    };
  } catch (error) {
    const message =
      error?.response?.data?.msg ||
      error?.message ||
      '获取文件列表失败';

    console.error('获取文件列表失败:', message);

    return {
      success: false,
      message
    };
  }
};
