using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace BisinessLogic
{
    public class Program_
    {
        public Program_(Guid id, string name, List<HotKey> hotKeys)
        {
            this.id = id;
            this.name = name;
            HotKeys = hotKeys;
        }

        public Guid id { get; }
        public string name { get; }
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
