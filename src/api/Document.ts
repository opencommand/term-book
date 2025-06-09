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

// open file
export const openFileApi = async (filename: string) => {
  try {
    const response = await apiClient.post('/document/open', {
      filename
    })

    return response.data
  } catch (error) {
    console.error('打开文件失败:', error)
    throw error
  }
}

// run cell
export const runCellApi = async (input: string) => {
  try {
    const response = await apiClient.post('/run-cell', {
      input
    })
    console.log(response.data);

    return response.data
  } catch (error) {
    console.error('运行失败:', error)
    throw error
  }
}

// save file
export const saveFileApi = async (data: {
  author: string;
  filename: string;
  cells: {
    exec_time: string;
    input: string;
    input_time: string;
    output: string;
    output_time: string;
  }[];
}) => {
  try {
    const response = await apiClient.post('/document/save', data);
    console.log('保存成功:', response.data);

    return response.data;
  } catch (error) {
    console.error('保存失败:', error);
    throw error;
  }
};
