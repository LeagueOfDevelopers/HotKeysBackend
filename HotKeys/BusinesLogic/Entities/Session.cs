using System;
using System.Collections.Generic;
using System.Linq;

namespace HotKeysLibrary.Entities
{
    public class Session
    {
        public int Id { get; }
        private List<HotKey> hotKeys;
        private Random _random;

        string Start(int id, Program test_program) //пока id int, потому что не факт что будет Guid
        {
            _random = new Random();
            hotKeys = test_program.hotKeys.ToList();
            var question = CreateQuestion();
            return question;
        }

        string MoveToNextQuestion()
        {
            if (hotKeys.Count() > 0)
            {
                var question = CreateQuestion();
                return question;
            }
            else
                return string.Empty;
        }

        string CreateQuestion()
        {
            var index = _random.Next(hotKeys.Count()); 
            return hotKeys[index].Description;
        }

        bool CheckAnswer(string question, List<Key> answer)
        {
            var hotKey = hotKeys.FirstOrDefault(hk => hk.Description == question);
            var result = (hotKey.Combination.ToList() == answer);
            return result;
        }

    }
}
