# simple_set_collector.py
def main() -> None:
    seen: set[str] = set()

    print("文字列を入力してください（exit/quit で終了）")
    while True:
        try:
            s = input("> ").strip()
            if s.lower() in {"exit", "quit"}:
                print("終了します。")
                break

            if s in seen:
                print("⚠️  もうすでにあります")
            else:
                seen.add(s)
                print("✅  追加しました")
        except KeyboardInterrupt:
            print("\nCtrl-C により終了します。")
            break

if __name__ == "__main__":
    main()
