import sys

def ksa(key):
    """key scheduling"""
    S = list(range(N))
    j = 0
    for i in range(0, N):
        j = (j + S[i] + ord(key[i % len(key)])) % N
        S[i], S[j] = S[j], S[i] # swap
    return S

def prga(S, n):
    """pseudo random generation"""
    i = 0
    j = 0
    keystream = []
    for i in range(0, n):
        i = (i + 1) % N
        j = (j + S[i]) % N
        S[i], S[j] = S[j], S[i]
        K = S[(S[i] + S[j]) % N]
        keystream.append(K)
    return keystream

def rc4(key, plaintext):
    S = ksa(key)
    return prga(S, len(plaintext))

def main():
    global N
    N = 256 # 2^8
    if len(sys.argv) < 3:
        raise Exception(f"{sys.argv[0]} <key> <value>")
    ciphertext = rc4(sys.argv[1], sys.argv[2])
    for index, char in enumerate(sys.argv[2]):
        print("%02x" % (ord(char) ^ ciphertext[index]), end="")
    print("\n", end="")

if __name__ == '__main__':
    main()
