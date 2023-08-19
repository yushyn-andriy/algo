from concurrent.futures import ProcessPoolExecutor, ThreadPoolExecutor, as_completed
import multiprocessing


import sys
import os



def search(file_path, exp):
    results = []
    try:
        with open(file_path, 'r', encoding='utf-8', errors='ignore') as f:
            for i, line in enumerate(f.readlines()):
                if exp in line:
                    results.append((file_path, i, line))
    except:
        return results
    return results


def get_paths(root_dir):
    file_paths = []
    for dirpath, dirnames, filenames in os.walk(root_dir):
        for filename in filenames:
            full_path = os.path.join(dirpath, filename)
            file_paths.append(full_path)
    return file_paths


def main(directory, exp):
    paths = get_paths(directory)
    futures = []
    with ProcessPoolExecutor(max_workers=4) as pool:
        for path in paths:
            future = pool.submit(search, path, exp)
            futures.append(future)
    
    for result in as_completed(futures):
        res = result.result()
        if res:
            for summary in res:
                print(summary)


if __name__ == '__main__':
    directory, exp = sys.argv[1], sys.argv[2]
    print(exp)
    main(directory, exp)
