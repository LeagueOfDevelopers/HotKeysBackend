using System;
using System.Collections.Generic;
using System.Linq;

namespace HotKeysLibrary.Entities
{
    public class Program
    {
        public Program(Guid id, string name, List<HotKey> hotKeys)
        {
            this.Id = id;
            this.Name = name;
            HotKeys = hotKeys;
        }

        public Guid Id { get; }
        public string Name { get; }
        private List<HotKey> HotKeys;
        public IEnumerable<HotKey> hotKeys => HotKeys.AsEnumerable();

        public void AddNewHotKey(HotKey hotKey)
        {
            HotKeys.Add(hotKey);
        }

        public void ChangeHotKey(Guid hotKeyId, HotKey newHotKey)
        {
            var id = HotKeys.FindIndex(hotk => hotk.Id == hotKeyId);
            HotKeys[id] = newHotKey;
        }
    }
}
