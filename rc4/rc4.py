import sys

def rc4(key, plaintext):
    return key + plaintext

def main():
    if len(sys.argv) < 3:
        print(f"{sys.argv[0]} <key> <plaintext>")
    ciphertext = rc4(sys.argv[1], sys.argv[2])
    print(ciphertext)

if __name__ == '__main__':
    main()
