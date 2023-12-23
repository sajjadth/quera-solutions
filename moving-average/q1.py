import numpy as np

def moving_average(data_list, window_size):
    res = []

    for i, _ in enumerate(data_list):
        if i + window_size > len(data_list):
            break
        window_data = np.array(data_list[i:i+window_size])
        res.append(round(np.mean(window_data), 2))

    return np.array(res)
