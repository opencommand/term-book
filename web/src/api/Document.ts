import axios from 'axios';
import config from './config';

const apiClient = axios.create({
  baseURL: config.API_BASE_URL,
  headers: { 'Content-Type': 'application/json' }
});

interface Response<T> {
  msg: string;
  data: T;
}

type DocumentGetResponse = Response<string[]>

// get fill list
export const getFileListApi = async () => {
  try {
    const res = ((await apiClient.get('/document/list')).data) as DocumentGetResponse;

    if (!Array.isArray(res.data)) {
      throw new Error('返回数据格式错误');
    }

    return {
      success: true,
      data: res.data,
      message: res.msg || '获取成功'
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

type RunCellResponse = Response<{
  output: string;
  output_time: string;
  exec_time: string | number;
}>

// run cell
export const runCellApi = async (input: string) => {
  try {
    const res = (await apiClient.post('/run-cell', {
      input
    })).data as RunCellResponse

    return res
  } catch (error) {
    console.error('运行失败:', error)
    throw error
  }
}

interface SaveFileRequest {
  author: string;
  filename: string;
  cells: {
    exec_time: string;
    input: string;
    input_time: string;
    output: string;
    output_time: string;
  }[];
}

// save file
export const saveFileApi = async (data: SaveFileRequest) => {
  try {
    const response = await apiClient.post('/document/save', data);
    console.log('保存成功:', response.data);

    return response.data;
  } catch (error) {
    console.error('保存失败:', error);
    throw error;
  }
};
